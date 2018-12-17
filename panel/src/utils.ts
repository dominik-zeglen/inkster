import { stringify as stringifyQs } from "qs";

export function maybe<T>(exp: () => T, d?: T) {
  try {
    const result = exp();
    return result === undefined ? d : result;
  } catch {
    return d;
  }
}

export function only<T>(obj: T, key: keyof T): boolean {
  return Object.keys(obj).every(
    objKey =>
      objKey === key ? obj[key] !== undefined : obj[key] === undefined,
  );
}

export function empty(obj: object): boolean {
  return Object.keys(obj).every(key => obj[key] === undefined);
}

export function renderCollection<T>(
  collection: T[],
  renderItem: (
    item: T | undefined,
    index: number | undefined,
    collection: T[],
  ) => any,
  renderEmpty?: (collection: T[]) => any,
) {
  if (collection === undefined) {
    return renderItem(undefined, undefined, collection);
  }
  if (collection.length === 0) {
    return !!renderEmpty ? renderEmpty(collection) : null;
  }
  return collection.map(renderItem);
}

export function mergeQs<T extends {} = {}>(qs: T, params: Partial<T>): string {
  return "?" + stringifyQs(Object.assign({}, qs, params));
}
