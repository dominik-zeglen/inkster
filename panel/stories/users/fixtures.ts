import * as faker from "faker";
import * as _ from "lodash";

faker.seed(4096);

function genUser() {
  const email = faker.internet.exampleEmail();
  const createdAt = faker.date.between(
    "2017-05-30T01:10:26+00:00",
    "2018-08-30T01:10:26+00:00"
  );
  return {
    id: "user" + faker.random.number(),
    email,
    createdAt,
    updatedAt: faker.random.boolean
      ? createdAt
      : faker.date.between(
          "2017-05-30T01:10:26+00:00",
          "2018-08-30T01:10:26+00:00"
        )
  };
}

export const users = _.times(10, genUser)
export const user = genUser();
