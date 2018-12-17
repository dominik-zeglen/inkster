import * as urlJoin from "url-join";

const home = "/";
const directorySection = "/directories/";
const pageSection = "/pages/";
const userSection = "/users/";

export const paths = {
  directoryList: directorySection,
  directoryDetails: (id: string) => urlJoin(directorySection, id),
  home,
  pageCreate: (id: string) => urlJoin(directorySection, id, "createPage"),
  pageDetails: (id: string) => urlJoin(pageSection, id),
  userList: userSection,
  userDetails: (id: string) => urlJoin(userSection, id),
  passwordRecovery: "/recover-password/",
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
};
export default urls;
