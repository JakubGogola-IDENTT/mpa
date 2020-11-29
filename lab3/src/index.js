/* eslint-disable no-console */
import { DelaunayRecursive } from './algorithms/recursive/DelaunayRecursive.js';
import { DelaunayIterative } from './algorithms/iterative/DelaunayIterative.js';
import { Vertex } from './algorithms/iterative/Vertex.js';
import { generatePoints } from './analytics/generatePoints.js';

const points = generatePoints(100);

const recursive = DelaunayRecursive.create(points);

const resRec = recursive.triangulate();

console.log(resRec);

const resIt = DelaunayIterative.create(
    points.map(point => Vertex.fromArray(point))
);

console.log(resIt);
