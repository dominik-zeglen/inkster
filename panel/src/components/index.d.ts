type Partial<T> = {
  [P in keyof T]?: T[P];
}

export type StandardProps = Partial<{
  className: string | undefined,
  style: React.CSSProperties | undefined,
  onClick: () => void
}>;
