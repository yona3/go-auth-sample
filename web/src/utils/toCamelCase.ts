export const toCamelCase = (str: string) => {
  return str.replace(/[-_](\w)/g, (match, letter) => {
    return letter.toUpperCase();
  });
};

export const toCamelCaseObjectKey = (obj: { [key: string]: unknown }) => {
  const newObj = {};
  for (const key in obj) {
    if (!Object.prototype.hasOwnProperty.call(obj, key)) return;
    (newObj as any)[toCamelCase(key)] = obj[key];
  }
  return newObj;
};
