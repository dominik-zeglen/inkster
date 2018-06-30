import * as faker from "faker";
import * as _ from "lodash";

faker.seed(4096);

const types = ["text", "longText"];
function genField() {
  const name = faker.commerce.department();
  return {
    id: "pagefield" + faker.random.number(),
    name,
    slug: faker.helpers.slugify(name),
    type: types[faker.random.number() % types.length],
    value: faker.lorem.sentence()
  };
}
function genDirectory() {
  return {
    id: "cat" + faker.random.number(),
    name: faker.commerce.department()
  };
}
function genPage() {
  const name = faker.commerce.department();
  return {
    id: "page" + faker.random.number(),
    name,
    slug: faker.helpers.slugify(name),
    parent: genDirectory(),
    fields: _.times((faker.random.number() % 6) + 1, genField)
  };
}

export const page = genPage();
