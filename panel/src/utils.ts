export function urlize(url: string) {
  return encodeURIComponent(url);
}
export function unurlize(url: string) {
  return decodeURIComponent(url);
}
