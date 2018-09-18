import * as React from "react";
import { Query } from "react-apollo";

import qUser from "../queries/qUser";
import UserDetailsPage from "../components/UserDetailsPage";
import Navigator from "../../components/Navigator";

interface Props {
  id: string;
}

export const UserDetails: React.StatelessComponent<Props> = ({ id }) => (
  <Navigator>
    {navigate => (
      <Query query={qUser} variables={{ id }}>
        {userData => (
          <UserDetailsPage
            disabled={userData.loading}
            loading={userData.loading}
            onBack={() => navigate("/users/")}
            onDelete={() => undefined}
            title={
              userData.data && userData.data.user
                ? userData.data.user.email
                : undefined
            }
            transaction={userData.loading ? "loading" : "default"}
            user={userData.data ? userData.data.user : undefined}
            onSubmit={() => undefined}
          />
        )}
      </Query>
    )}
  </Navigator>
);
export default UserDetails;
