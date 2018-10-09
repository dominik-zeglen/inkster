import * as React from "react";
import { withRouter } from "react-router";

import { User, UserContext } from "./auth/components/AuthProvider";
import AppLayout from "./components/AppLayout";

export const AppRoot = withRouter(({ children, history, location }) => {
  const section = location.pathname.split("/")[1] || "home";
  const handleSectionClick = (sectionName: string) => () => {
    switch (sectionName) {
      case "home":
        history.push("/");
        break;
      case "directories":
        history.push("/directories");
        break;
      case "users":
        history.push("/users");
        break;
    }
  };
  return (
    <UserContext.Consumer>
      {({ user, logout }) => (
        <AppLayout
          section={section}
          user={user as User}
          onLogout={logout}
          onSectionClick={handleSectionClick}
        >
          {children}
        </AppLayout>
      )}
    </UserContext.Consumer>
  );
});
export default AppRoot;
