/* eslint-disable class-methods-use-this */
export class Delaunay {
    /**
     * @param {string} name
     */
    constructor(name) {
        this.name = name;
    }

    triangulate() {
        throw new Error('Method "triangulate" has to be implemented.');
    }

    /**
     * @returns {Delaunay}
     */
    static create() {
        return new Delaunay();
    }
}
