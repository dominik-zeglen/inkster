import * as faker from 'faker'
import * as _ from 'lodash'

faker.seed(4096)

function genPage() {
  const name = faker.lorem.word();
  return {
    id: "page" + faker.random.number(),
    name,
    slug: faker.helpers.slugify(name)
  }
}
function genDirectory() {
  return {
    id: "cat" + faker.random.number(),
    name: faker.commerce.department(),
    pages: _.times(faker.random.number() % 10, genPage)
  }
}

export const directories = _.times(10, genDirectory)
export const directory = genDirectory()
