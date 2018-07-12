export function urlize(url: string) {
  return url.replace(/\//g, "%2F").replace(/\+/g, "%2B");
}
export function unurlize(url: string) {
  return url.replace(/%2F/g, "/").replace(/%2B/g, "+");
}
