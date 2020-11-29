export class Vertex {
    /**
     * @param {number} x
     * @param {number} y
     */
    constructor(x, y) {
        this.x = x;
        this.y = y;
    }

    /**
     * @param {Vertex} vertex
     * @returns {boolean}
     */
    equals({ x, y }) {
        return this.x === x && this.y === y;
    }

    /**
     * @param {number[]} point
     * @returns {Vertex}
     */
    static fromArray([x, y]) {
        return new Vertex(x, y);
    }
}
