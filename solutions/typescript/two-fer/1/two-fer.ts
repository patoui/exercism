export function twoFer(name?: string): string {
  name ??= 'you';
  return `One for ${name}, one for me.`;
}
