/* eslint-disable class-methods-use-this */
import { Vertex } from './Vertex.js';
import { Edge } from './Edge.js';
import { Triangle } from './Triangle.js';
import { Delaunay } from '../Delaunay.js';

export class DelaunayIterative extends Delaunay {
    /**
     * @param {Vertex[]} points
     */
    constructor(points) {
        super('iterative');
        this.vertices = points.map(point => Vertex.fromArray(point));
        this.iterations = 0;
    }

    /**
     *
     * @param {Vertex[]} vertices
     * @returns {DelaunayIterative}
     */
    static create(vertices) {
        return new DelaunayIterative(vertices);
    }

    /**
     * @returns {Triangle}
     */
    superTriangle() {
        let minY = Infinity;
        let minX = minY;
        let maxY = -Infinity;
        let maxX = maxY;

        this.vertices.forEach(vertex => {
            this.iterations++;
            minX = Math.min(minX, vertex.x);
            minY = Math.min(minX, vertex.y);
            maxX = Math.max(maxX, vertex.x);
            maxY = Math.max(maxX, vertex.y);
        });

        const dx = (maxX - minX) * 10;
        const dy = (maxY - minY) * 10;

        const v0 = new Vertex(minX - dx, minY - dy * 3);
        const v1 = new Vertex(minX - dx, maxY + dy);
        const v2 = new Vertex(maxX + dx * 3, maxY + dy);

        return new Triangle(v0, v1, v2);
    }

    /**
     * @param {Vertex} vertex
     * @param {Triangle[]} triangles
     * @returns {Triangle[]}
     */
    addVertex(vertex, triangles) {
        const edges = [];

        // Remove triangles with circumcircles containing the vertex
        const trianglesFiltered = triangles.filter(triangle => {
            this.iterations++;
            if (triangle.inCircumcircle(vertex)) {
                edges.push(new Edge(triangle.v0, triangle.v1));
                edges.push(new Edge(triangle.v1, triangle.v2));
                edges.push(new Edge(triangle.v2, triangle.v0));
                return false;
            }
            return true;
        });

        // Get unique edges
        const uqEdges = this.uniqueEdges(edges);

        // Create new triangles from the unique edges and new vertex
        uqEdges.forEach(edge => {
            this.iterations++;
            trianglesFiltered.push(new Triangle(edge.v0, edge.v1, vertex));
        });

        return trianglesFiltered;
    }

    /**
     * @param {Edge[]} edges
     * @returns {Edge[]}
     */
    uniqueEdges(edges) {
        const uniqueEdges = [];
        for (let i = 0; i < edges.length; ++i) {
            this.iterations++;
            let isUnique = true;

            // See if edge is unique
            for (let j = 0; j < edges.length; ++j) {
                this.iterations++;
                if (i !== j && edges[i].equals(edges[j])) {
                    isUnique = false;
                    break;
                }
            }

            // Edge is unique, add to unique edges array
            if (isUnique) {
                uniqueEdges.push(edges[i]);
            }
        }

        return uniqueEdges;
    }

    triangulate() {
        const st = this.superTriangle();

        let triangles = [st];

        this.vertices.forEach(vertex => {
            this.iterations++;
            triangles = this.addVertex(vertex, triangles);
        });

        triangles = triangles.filter(
            triangle =>
                !(
                    triangle.v0 === st.v0 ||
                    triangle.v0 === st.v1 ||
                    triangle.v0 === st.v2 ||
                    triangle.v1 === st.v0 ||
                    triangle.v1 === st.v1 ||
                    triangle.v1 === st.v2 ||
                    triangle.v2 === st.v0 ||
                    triangle.v2 === st.v1 ||
                    triangle.v2 === st.v2
                )
        );

        return [triangles, this.iterations];
    }
}
