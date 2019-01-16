declare module "react-feather" {
  import { StatelessComponent } from "react";
  import * as ReactFeather from "react-feather";

  export interface FeatherIcon {
    size?: number;
    color?: string;
  }

  export interface ReactFeatherModule {
    [key: string]: StatelessComponent<FeatherIcon>;
  }

  export const AlertTriangle: StatelessComponent<FeatherIcon>;
  export const ArrowLeft: StatelessComponent<FeatherIcon>;
  export const ArrowRight: StatelessComponent<FeatherIcon>;
  export const Bold: StatelessComponent<FeatherIcon>;
  export const Box: StatelessComponent<FeatherIcon>;
  export const ChevronDown: StatelessComponent<FeatherIcon>;
  export const ChevronLeft: StatelessComponent<FeatherIcon>;
  export const ChevronUp: StatelessComponent<FeatherIcon>;
  export const File: StatelessComponent<FeatherIcon>;
  export const FileText: StatelessComponent<FeatherIcon>;
  export const Folder: StatelessComponent<FeatherIcon>;
  export const Home: StatelessComponent<FeatherIcon>;
  export const Image: StatelessComponent<FeatherIcon>;
  export const Info: StatelessComponent<FeatherIcon>;
  export const LogOut: StatelessComponent<FeatherIcon>;
  export const Maximize2: StatelessComponent<FeatherIcon>;
  export const Minimize2: StatelessComponent<FeatherIcon>;
  export const Plus: StatelessComponent<FeatherIcon>;
  export const Settings: StatelessComponent<FeatherIcon>;
  export const Trash: StatelessComponent<FeatherIcon>;
  export const User: StatelessComponent<FeatherIcon>;
  export const Users: StatelessComponent<FeatherIcon>;
  export const X: StatelessComponent<FeatherIcon>;
}
