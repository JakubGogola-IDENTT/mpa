import { Vertex } from './Vertex.js';

export class Triangle {
    /**
     * @param {Vertex} v0
     * @param {Vertex} v1
     * @param {Vertex} v2
     */
    constructor(v0, v1, v2) {
        this.v0 = v0;
        this.v1 = v1;
        this.v2 = v2;

        this.radius = NaN;
        this.center = null;
    }

    calcCircumcircle() {
        const A = this.v1.x - this.v0.x;
        const B = this.v1.y - this.v0.y;
        const C = this.v2.x - this.v0.x;
        const D = this.v2.y - this.v0.y;

        const E = A * (this.v0.x + this.v1.x) + B * (this.v0.y + this.v1.y);
        const F = C * (this.v0.x + this.v2.x) + D * (this.v0.y + this.v2.y);

        const G =
            2.0 * (A * (this.v2.y - this.v1.y) - B * (this.v2.x - this.v1.x));

        let dx;
        let dy;

        if (Math.round(Math.abs(G)) === 0) {
            const minx = Math.min(this.v0.x, this.v1.x, this.v2.x);
            const miny = Math.min(this.v0.y, this.v1.y, this.v2.y);
            const maxx = Math.max(this.v0.x, this.v1.x, this.v2.x);
            const maxy = Math.max(this.v0.y, this.v1.y, this.v2.y);

            this.center = new Vertex((minx + maxx) / 2, (miny + maxy) / 2);

            dx = this.center.x - minx;
            dy = this.center.y - miny;
        } else {
            const cx = (D * E - B * F) / G;
            const cy = (A * F - C * E) / G;

            this.center = new Vertex(cx, cy);

            dx = this.center.x - this.v0.x;
            dy = this.center.y - this.v0.y;
        }

        this.radius = Math.sqrt(dx * dx + dy * dy);
    }

    /**
     * @param {Vertex} vertex
     * @returns {boolean}
     */
    inCircumcircle({ x, y }) {
        if (!this.center) {
            throw new Error('Center is not defined.');
        }

        const dx = this.center.x - x;
        const dy = this.center.y - y;

        return Math.sqrt(dx * dx + dy * dy) <= this.radius;
    }
}
