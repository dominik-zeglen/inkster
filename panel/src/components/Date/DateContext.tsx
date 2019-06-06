import * as React from "react";

const DateContext = React.createContext<number>(0);

export const { Consumer, Provider } = DateContext;
export default DateContext;
