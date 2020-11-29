import { Vertex } from './Vertex.js';

export class Edge {
    /**
     * @param {Vertex} v0
     * @param {Vertex} v1
     */
    constructor(v0, v1) {
        this.v0 = v0;
        this.v1 = v1;
    }

    /**
     * @param {Edge} edge
     * @returns {boolean}
     */
    equals({ v0, v1 }) {
        return (
            (this.v0.equals(v0) && this.v1.equals(v1)) ||
            (this.v0.equals(v1) && this.v1.equals(v0))
        );
    }

    inverse() {
        return new Edge(this.v1, this.v0);
    }
}
