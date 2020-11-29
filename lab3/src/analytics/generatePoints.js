import MersenneTwister from 'mersenne-twister';

const generator = new MersenneTwister();

const randomPoint = (bound = 100) => {
    const mult = generator.random_int() % bound;
    const sign = mult % 2;

    const value = generator.random();

    return sign === 1 ? -value : value;
};

export const generatePoints = (size, bound = 100) => {
    return new Array(size)
        .fill(0)
        .map(() => [randomPoint(bound), randomPoint(bound)]);
};
