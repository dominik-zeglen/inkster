declare module "draft-js-plugins-editor" {
  import * as DraftJSPluginsEditor from "draft-js-plugins-editor";
  import { EditorProps, EditorState } from "draft-js";
  import { Component } from "react";

  export interface PluginsEditorProps extends EditorProps {
    plugins: any[];
  }

  export const createEditorStateWithText: (text: string) => any;
  export const Editor: (
    props: PluginsEditorProps
  ) => Component<PluginsEditorProps, EditorState>;
  export const composeDecorators: any;
  export default Editor;
}
