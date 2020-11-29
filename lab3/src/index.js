/* eslint-disable no-console */
import { DelaunayRecursive } from './algorithms/recursive/DelaunayRecursive.js';
import { DelaunayIterative } from './algorithms/iterative/DelaunayIterative.js';
import { testing } from './analytics/testing.js';

testing(DelaunayRecursive, './recursive.txt');
testing(DelaunayIterative, './iterative.txt');
