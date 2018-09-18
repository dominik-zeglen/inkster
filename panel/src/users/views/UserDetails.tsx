import * as React from "react";
import { Mutation, Query } from "react-apollo";

import mUpdateUser from "../queries/mUpdateUser";
import mRemoveUser from "../queries/mRemoveUser";
import qUser from "../queries/qUser";
import UserDetailsPage from "../components/UserDetailsPage";
import Navigator from "../../components/Navigator";

interface Props {
  id: string;
}

export const UserDetails: React.StatelessComponent<Props> = ({ id }) => (
  <Navigator>
    {navigate => {
      const handleRemoveUser = () => navigate("/users/");
      return (
        <Query query={qUser} variables={{ id }}>
          {userData => (
            <Mutation mutation={mRemoveUser} onCompleted={handleRemoveUser}>
              {removeUser => (
                <Mutation mutation={mUpdateUser}>
                  {updateUser => (
                    <UserDetailsPage
                      disabled={userData.loading}
                      loading={userData.loading}
                      onBack={() => navigate("/users/")}
                      onDelete={() => removeUser({ variables: { id } })}
                      title={
                        userData.data && userData.data.user
                          ? userData.data.user.email
                          : undefined
                      }
                      transaction={userData.loading ? "loading" : "default"}
                      user={userData.data ? userData.data.user : undefined}
                      onSubmit={variables => updateUser({variables: {id, input: variables}})}
                    />
                  )}
                </Mutation>
              )}
            </Mutation>
          )}
        </Query>
      );
    }}
  </Navigator>
);
export default UserDetails;
