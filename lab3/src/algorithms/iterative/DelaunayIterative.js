/* eslint-disable class-methods-use-this */
import { Vertex } from './Vertex.js';
import { Edge } from './Edge.js';
import { Triangle } from './Triangle.js';

export class DelaunayIterative {
    /**
     * @param {Vertex[]} vertices
     */
    constructor(vertices) {
        this.vertices = vertices;
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
            let isUnique = true;

            // See if edge is unique
            for (let j = 0; j < edges.length; ++j) {
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
        const superTriangle = this.superTriangle();

        let triangles = [superTriangle];

        this.vertices.forEach(vertex => {
            triangles.addVertex(vertex, triangles);
        });

        const { v0: sv0, v1: sv1, v2: sv2 } = superTriangle;
        const stvs = [sv0, sv1, sv2];

        triangles = triangles.filter(
            ({ v0, v1, v2 }) => ![v0, v1, v2].some(v => stvs.includes(v))
        );

        return triangles;
    }
}
