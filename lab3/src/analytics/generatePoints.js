import MersenneTwister from 'mersenne-twister';

const generator = new MersenneTwister();

const randomPoint = () => {
    const mult = generator.random_int();
    const sign = mult % 2;

    const value = generator.random();

    return sign === 1 ? -value : value;
};

export const generatePoints = size => {
    return new Array(size).fill(0).map(() => [randomPoint(), randomPoint()]);
};
