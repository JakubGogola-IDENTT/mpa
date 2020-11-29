const ccw = ([ax, ay], [bx, by], [cx, cy]) =>
    (bx - ax) * (cy - ay) - (by - ay) * (cx - ax) > 0;

const rightOf = (x, { dest, orig }) => ccw(x, dest, orig);

const leftOf = (x, { dest, orig }) => ccw(x, orig, dest);

const valid = ({ dest }, basel) => rightOf(dest, basel);

const inCircle = (a, b, c, d) => {
    if (
        (a[0] === d[0] && a[1] === d[1]) ||
        (b[0] === d[0] && b[1] === d[1]) ||
        (c[0] === d[0] && c[1] === d[1])
    ) {
        return false;
    }

    const sa = a[0] * a[0] + a[1] * a[1];
    const sb = b[0] * b[0] + b[1] * b[1];
    const sc = c[0] * c[0] + c[1] * c[1];
    const sd = d[0] * d[0] + d[1] * d[1];

    const d1 = sc - sd;
    const d2 = c[1] - d[1];
    const d3 = c[1] * sd - sc * d[1];
    const d4 = c[0] - d[0];
    const d5 = c[0] * sd - sc * d[0];
    const d6 = c[0] * d[1] - c[1] * d[0];

    return (
        a[0] * (b[1] * d1 - sb * d2 + d3) -
            a[1] * (b[0] * d1 - sb * d4 + d5) +
            sa * (b[0] * d2 - b[1] * d4 + d6) -
            b[0] * d3 +
            b[1] * d5 -
            sb * d6 >
        1
    );
};

class QuadEdge {
    constructor(onext, rot, orig) {
        this.onext = onext;
        this.rot = rot;
        this.orig = orig;
        this.mark = false;
    }

    get sym() {
        return this.rot.rot;
    }

    get dest() {
        return this.sym.orig;
    }

    get rotSym() {
        return this.rot.sym;
    }

    get oprev() {
        return this.rot.onext.rot;
    }

    get dprev() {
        return this.rotSym.onext.rotSym;
    }

    get lnext() {
        return this.rotSym.onext.rot;
    }

    get lprev() {
        return this.onext.sym;
    }

    get rprev() {
        return this.sym.onext;
    }
}

const makeEdge = (orig, dest) => {
    const q0 = new QuadEdge(null, null, orig);
    const q1 = new QuadEdge(null, null, null);
    const q2 = new QuadEdge(null, null, dest);
    const q3 = new QuadEdge(null, null, null);

    // create the segment
    q0.onext = q0;
    q2.onext = q2; // lonely segment: no "next" quadedge
    q1.onext = q3;
    q3.onext = q1; // in the dual: 2 communicating facets

    // dual switch
    q0.rot = q1;
    q1.rot = q2;
    q2.rot = q3;
    q3.rot = q0;

    return q0;
};

const splice = (a, b) => {
    const alpha = a.onext.rot;
    const beta = b.onext.rot;

    const t2 = a.onext;
    const t3 = beta.onext;
    const t4 = alpha.onext;

    /* eslint-disable no-param-reassign */
    a.onext = b.onext;
    b.onext = t2;
    /* eslint-enable no-param-reassign */
    alpha.onext = t3;
    beta.onext = t4;
};

const connect = (a, b) => {
    const q = makeEdge(a.dest, b.orig);
    splice(q, a.lnext);
    splice(q.sym, b);
    return q;
};

const deleteEdge = q => {
    splice(q, q.oprev);
    splice(q.sym, q.sym.oprev);
};

const delaunay = s => {
    let a;
    let b;
    let c;
    let t;

    if (s.length === 2) {
        a = makeEdge(s[0], s[1]);
        return {
            le: a,
            re: a.sym,
        };
    }
    if (s.length === 3) {
        a = makeEdge(s[0], s[1]);
        b = makeEdge(s[1], s[2]);
        splice(a.sym, b);

        if (ccw(s[0], s[1], s[2])) {
            c = connect(b, a);
            return {
                le: a,
                re: b.sym,
            };
        }
        if (ccw(s[0], s[2], s[1])) {
            c = connect(b, a);
            return {
                le: c.sym,
                re: c,
            };
        } // All three points are colinear
        return {
            le: a,
            re: b.sym,
        };
    } // |S| >= 4
    const halfLength = Math.ceil(s.length / 2);
    const left = delaunay(s.slice(0, halfLength));
    const right = delaunay(s.slice(halfLength));

    let ldo = left.le;
    let ldi = left.re;
    let rdi = right.le;
    let rdo = right.re;

    // Compute the lower common tangent of L and R
    do {
        if (leftOf(rdi.orig, ldi)) ldi = ldi.lnext;
        else if (rightOf(ldi.orig, rdi)) rdi = rdi.rprev;
        else break;
        // eslint-disable-next-line no-constant-condition
    } while (true);

    let basel = connect(rdi.sym, ldi);
    if (ldi.orig === ldo.orig) ldo = basel.sym;
    if (rdi.orig === rdo.orig) rdo = basel;

    // This is the merge loop.
    do {
        // Locate the first L point (lcand.Dest) to be encountered by the rising bubble,
        // and delete L edges out of base1.Dest that fail the circle test.
        let lcand = basel.sym.onext;
        if (valid(lcand, basel)) {
            while (
                inCircle(basel.dest, basel.orig, lcand.dest, lcand.onext.dest)
            ) {
                t = lcand.onext;
                deleteEdge(lcand);
                lcand = t;
            }
        }

        // Symmetrically, locate the first R point to be hit, and delete R edges
        let rcand = basel.oprev;
        if (valid(rcand, basel)) {
            while (
                inCircle(basel.dest, basel.orig, rcand.dest, rcand.oprev.dest)
            ) {
                t = rcand.oprev;
                deleteEdge(rcand);
                rcand = t;
            }
        }

        // If both lcand and rcand are invalid, then basel is the upper common tangent
        if (!valid(lcand, basel) && !valid(rcand, basel)) break;

        // The next cross edge is to be connected to either lcand.Dest or rcand.Dest
        // If both are valid, then choose the appropriate one using the InCircle test
        if (
            !valid(lcand, basel) ||
            (valid(rcand, basel) &&
                inCircle(lcand.dest, lcand.orig, rcand.orig, rcand.dest))
        )
            // Add cross edge basel from rcand.Dest to basel.Dest
            basel = connect(rcand, basel.sym);
        // Add cross edge base1 from basel.Org to lcand. Dest
        else basel = connect(basel.sym, lcand.sym);
        // eslint-disable-next-line no-constant-condition
    } while (true);

    return {
        le: ldo,
        re: rdo,
    };
};

export class DelaunayRecursive {
    /**
     * @param {Array<number>[]} points
     */
    constructor(points = []) {
        this.points = points;
    }

    static create(points = []) {
        return new DelaunayRecursive(points);
    }

    /**
     * @returns {Array<number>[]}
     */
    triangulate() {
        const pts = this.points;

        // Initial sorting of the points
        pts.sort((a, b) => {
            if (a[0] === b[0]) return a[1] - b[1];
            return a[0] - b[0];
        });

        // Remove duplicates
        for (let i = pts.length - 1; i >= 1; i--)
            if (pts[i][0] === pts[i - 1][0] && pts[i][1] === pts[i - 1][1])
                pts.splice(i, 1); // Costly operation, but there shouldn't be that many duplicates

        if (pts.length < 2) return [];

        let quadEdge = delaunay(pts).le;

        // All edges marked false
        const faces = [];
        let queueIndex = 0;
        const queue = [quadEdge];

        // Mark all outer edges as visited
        while (leftOf(quadEdge.onext.dest, quadEdge)) quadEdge = quadEdge.onext;

        let curr = quadEdge;
        do {
            queue.push(curr.sym);
            curr.mark = true;
            curr = curr.lnext;
        } while (curr !== quadEdge);

        do {
            const edge = queue[queueIndex++];
            if (!edge.mark) {
                // Stores the edges for a visited triangle. Also pushes sym (neighbour) edges on stack to visit later.
                curr = edge;
                do {
                    faces.push(curr.orig);
                    if (!curr.sym.mark) queue.push(curr.sym);

                    curr.mark = true;
                    curr = curr.lnext;
                } while (curr !== edge);
            }
        } while (queueIndex < queue.length);

        return faces;
    }
}
