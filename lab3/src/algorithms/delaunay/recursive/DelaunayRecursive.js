export class DelaunayRecursive {
    /**
     * @param {Array<{x: number, y: number}>} points
     */
    constructor(points = []) {
        this.points = points;
    }

    static create(points = []) {
        return new DelaunayRecursive(points);
    }

    /**
     * @returns {Array<{x: number, y: number}>}
     */
    triangulate() {
        let { points } = this;

        // Sorting points
        points.sort((p1, p2) => (p1.x === p2.x ? p1.y - p2.y : p1.x - p2.x));

        // Removing duplicates
        // TODO: may not work
        points = points.filter((pt, idx, self) => {
            const succ = self[idx + 1];

            if (!succ) {
                return true;
            }

            return pt.x !== succ.x || pt.y !== succ.y;
        });

        if (points.length < 2) {
            return [];
        }

        returns []
    }
}
