import * as urlJoin from "url-join";

export const home = "/";
export const directorySection = "/directories/";
export const pageSection = "/pages/";
export const userSection = "/users/";
export const websiteSettingsSection = "/website/";

export const paths = {
  directoryList: directorySection,
  directoryDetails: (id: string) => urlJoin(directorySection, id),
  home,
  pageCreate: (id: string) => urlJoin(directorySection, id, "createPage"),
  pageDetails: (id: string) => urlJoin(pageSection, id),
  userList: userSection,
  userDetails: (id: string) => urlJoin(userSection, id),
  passwordRecovery: "/recover-password/",
  websiteSettings: websiteSettingsSection,
};

export const urls = {
  directoryList: paths.directoryList,
  directoryDetails: (id: string) =>
    paths.directoryDetails(encodeURIComponent(id)),
  home,
  pageCreate: (id: string) => paths.pageCreate(encodeURIComponent(id)),
  pageDetails: (id: string) => paths.pageDetails(encodeURIComponent(id)),
  userList: paths.userList,
  userDetails: (id: string) => paths.userDetails(encodeURIComponent(id)),
  passwordRecovery: paths.passwordRecovery,
  websiteSettings: websiteSettingsSection,
};
export default urls;
