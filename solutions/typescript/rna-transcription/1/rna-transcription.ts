const map = {
  G: 'C',
  C: 'G',
  T: 'A',
  A: 'U'
} as const;

type Dna = keyof typeof map;

function isDna(char: string): char is Dna {
  return char in map;
}

export function toRna(nucleotides: string): string {
  let result = '';

  for (const char of nucleotides) {
    if (!isDna(char)) {
      throw new Error('Invalid input DNA.');
    }

    result += map[char];
  }

  return result;
}
