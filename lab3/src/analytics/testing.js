import fs, { WriteStream } from 'fs';
import { DelaunayRecursive } from '../algorithms/recursive/DelaunayRecursive.js';
import { DelaunayIterative } from '../algorithms/iterative/DelaunayIterative.js';
import { generatePoints } from './generatePoints.js';


/**
 * @param {Function} instance
 * @param {string} path
 * @param {number} lowerBound
 * @param {number} upperBound
 * @param {number} step
 * @param {number} repetitions
 */
export const testing = (
    instance,
    lowerBound = 100,
    upperBound = 5000,
    step = 100,
    repetitions = 500
) => {
    const ws = fs.createWriteStream('./results.txt');

    for (let i = lowerBound; i <= upperBound; i += step) {
        console.log(`Progress: ${i}`);

        for (let j = 0; j < repetitions; j++) {
            const points = generatePoints(i, 100);

            const algorithm = instance.create(points);

            const [, iterations] = algorithm.triangulate();

            ws.write(`${i},${j},${iterations}\n`);
        }
    }

    ws.end();
};
