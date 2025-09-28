export function decodedResistorValue(colors: string[]) {
  const result = colors.slice(0, 2).reduce(
    (carry: string, value: string) => carry + colorCode(value).toString(),
    ''
  );

  let num = Number(result) * Math.max(1, 10 ** colorCode(colors[2]));

  if (num >= 1_000_000_000) {
    num /= 1_000_000_000;
    return `${num} gigaohms`;
  }

  if (num >= 1_000_000) {
    num /= 1_000_000;
    return `${num} megaohms`;
  }

  if (num >= 1000) {
    num /= 1000;
    return `${num} kiloohms`;
  }

  return `${num} ohms`;
}

export const colorCode = (color: string): number => {
  return COLORS.findIndex(v => v === color);
}

export const COLORS = [
  'black',
  'brown',
  'red',
  'orange',
  'yellow',
  'green',
  'blue',
  'violet',
  'grey',
  'white',
];
