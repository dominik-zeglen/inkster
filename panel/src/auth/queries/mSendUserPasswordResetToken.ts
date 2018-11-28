import gql from "graphql-tag";

import { TypedMutation } from "../../api";
import {
  SendUserPasswordResetToken,
  SendUserPasswordResetTokenVariables,
} from "./types/SendUserPasswordResetToken";

const mSendUserPasswordResetToken = gql`
  mutation SendUserPasswordResetToken($email: String!) {
    sendUserPasswordResetToken(email: $email)
  }
`;

export default TypedMutation<
  SendUserPasswordResetToken,
  SendUserPasswordResetTokenVariables
>(mSendUserPasswordResetToken);
