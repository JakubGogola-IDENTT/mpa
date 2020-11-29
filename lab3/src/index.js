import { DelaunayRecursive } from './algorithms/DelaunayRecursive.js';

const recursive = DelaunayRecursive.create(
    [[1.1, 2.3],
    [4.5, 2.3],
    [10.1, 12.3]]
);

const result = recursive.triangulate();

console.log(result);
