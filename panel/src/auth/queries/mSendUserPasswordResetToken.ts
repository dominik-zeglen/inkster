import gql from "graphql-tag";

const mSendUserPasswordResetToken = gql`
  mutation SendUserPasswordResetToken($email: String!) {
    sendUserPasswordResetToken(email: $email)
  }
`;
export interface Variables {
  email: string;
}
export interface Result {
  sendUserPasswordResetToken: boolean;
}
export default mSendUserPasswordResetToken;
