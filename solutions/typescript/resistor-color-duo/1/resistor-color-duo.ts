export function decodedValue(colors: string[]): number {
  const result = colors.slice(0, 2).reduce(
    (carry: string, value: string) => carry + colorCode(value).toString(),
    ''
  );

  return Number(result);
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