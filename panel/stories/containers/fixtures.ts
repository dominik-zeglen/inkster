import * as faker from 'faker'
import * as _ from 'lodash'

faker.seed(4096)

function genContainer() {
  return {
    id: "cat" + faker.random.number(),
    name: faker.commerce.department()
  }
}

export const containers = _.times(10, genContainer)
export const container = genContainer()
