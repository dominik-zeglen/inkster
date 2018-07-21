import * as React from "react";
import { withRouter } from "react-router";

import AppLayout from "./components/AppLayout";

export const AppRoot = withRouter(({ children, history, location }) => {
  const section = location.pathname.split("/")[1];
  const handleSectionClick = (section: string) => () => {
    switch (section) {
      case "home":
        history.push("/");
        break;
      case "directories":
        history.push("/directories");
        break;
    }
  };
  return (
    <AppLayout section={section} onSectionClick={handleSectionClick}>
      {children}
    </AppLayout>
  );
});
export default AppRoot;
