/* eslint-disable no-console */
import { DelaunayRecursive } from './algorithms/recursive/DelaunayRecursive.js';
import { DelaunayIterative } from './algorithms/iterative/DelaunayIterative.js';
import { Vertex } from './algorithms/iterative/Vertex.js';

const points = [
    [1.1, 2.3],
    [4.5, 2.3],
    [10.1, 12.3],
    [11.2, 21.37],
    [6.9, 6.9],
];

const recursive = DelaunayRecursive.create(points);

const resRec = recursive.triangulate();

console.log(resRec);

const resIt = DelaunayIterative.create(
    points.map(point => Vertex.fromArray(point))
);

console.log(resIt);
