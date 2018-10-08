export const urls = {
  directoryDetails: (id?: string) => `/directories/${id ? encodeURIComponent(id) : ""}`,
  pageCreate: (id: string) => `/directories/${encodeURIComponent(id)}/createPage`,
  pageDetails: (id: string) => `/pages/${encodeURIComponent(id)}`,
  userDetails: (id: string) => `/users/${encodeURIComponent(id)}`,
  passwordRecovery: `/recover-password`
};
export default urls;
