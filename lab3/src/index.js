/* eslint-disable no-console */
import { DelaunayRecursive } from './algorithms/recursive/DelaunayRecursive.js';
import { DelaunayIterative } from './algorithms/iterative/DelaunayIterative.js';
import { Vertex } from './algorithms/iterative/Vertex.js';
import { generatePoints } from './analytics/generatePoints.js';
import { testing } from './analytics/testing.js';

testing(DelaunayRecursive);