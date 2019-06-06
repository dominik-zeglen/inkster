import * as React from "react";
import withStyles, { WithStyles } from "react-jss";
import { EditorState, convertFromRaw, convertToRaw, RichUtils } from "draft-js";
import DraftEditor, { composeDecorators } from "draft-js-plugins-editor";
import createInlineToolbarPlugin, {
  Separator,
} from "draft-js-inline-toolbar-plugin";
import {
  ItalicButton,
  BoldButton,
  UnderlineButton,
  HeadlineOneButton,
  HeadlineTwoButton,
  HeadlineThreeButton,
  UnorderedListButton,
  OrderedListButton,
  BlockquoteButton,
} from "draft-js-buttons";
import createImagePlugin from "draft-js-image-plugin";
import createAlignmentPlugin from "draft-js-alignment-plugin";
import createFocusPlugin from "draft-js-focus-plugin";
import createResizeablePlugin from "draft-js-resizeable-plugin";
import createBlockDndPlugin from "draft-js-drag-n-drop-plugin";
import createDragNDropUploadPlugin from "@mikeljames/draft-js-drag-n-drop-upload-plugin";
import InputFocus from "aurora-ui-kit/dist/components/InputFocus";
import InputLabel from "aurora-ui-kit/dist/components/InputLabel";
import Typography from "aurora-ui-kit/dist/components/Typography";

import i18n from "../i18n";

interface Props {
  id?: string;
  label?: string;
  helperText?: string;
  error?: boolean;
  name: string;
  initialValue?: string;
  onChange: (event: React.ChangeEvent<any>) => void;
}
interface State {
  editorState: any;
  focused: boolean;
}

// Plugins
const mockUpload = (opts: any) => console.log(opts);
const focusPlugin = createFocusPlugin();
const resizeablePlugin = createResizeablePlugin();
const blockDndPlugin = createBlockDndPlugin();
const alignmentPlugin = createAlignmentPlugin();
const { AlignmentTool } = alignmentPlugin;

const decorator = composeDecorators(
  resizeablePlugin.decorator,
  alignmentPlugin.decorator,
  focusPlugin.decorator,
  blockDndPlugin.decorator,
);
const imagePlugin = createImagePlugin({ decorator });

const dragNDropFileUploadPlugin = createDragNDropUploadPlugin({
  handleUpload: mockUpload,
  addImage: imagePlugin.addImage,
});

const decorate = withStyles(theme => ({
  editor: {
    background: theme.mixins.fade(theme.colors.primary.main, 0.05),
    border: `1px solid ${theme.colors.primary.lightest}`,
    marginBottom: theme.spacing,
    marginTop: 0,
    padding: theme.spacing,
    transition: theme.transition.time,
  },
  editorContainer: {
    paddingTop: theme.spacing * 3,
  },
  headlineButton: {
    background: "transparent",
    color: "#888",
    fontSize: 21,
    border: 0,
    verticalAlign: "bottom",
    position: "relative" as "relative",
    top: -1,
    height: 34,
    width: 36,
  },
  headlineButtonWrapper: {
    display: "inline-block" as "inline-block",
  },
  toolbar: {
    display: "flex" as "flex",
    position: "absolute" as "absolute",
  },
}));

class HeadlinesPicker extends React.Component<any, any> {
  componentDidMount() {
    setTimeout(() => {
      window.addEventListener("click", this.onWindowClick);
    });
  }

  componentWillUnmount() {
    window.removeEventListener("click", this.onWindowClick);
  }

  onWindowClick = () => this.props.onOverrideContent(undefined);

  render() {
    const buttons = [HeadlineOneButton, HeadlineTwoButton, HeadlineThreeButton];
    return (
      <div>
        {buttons.map((Button, i) => (
          <Button key={i} {...this.props} />
        ))}
      </div>
    );
  }
}

const HeadlinesButton = decorate(
  class HeadlinesButtonComponent extends React.Component<any, any> {
    onMouseDown = (event: any) => event.preventDefault();

    onClick = () => this.props.onOverrideContent(HeadlinesPicker);

    render() {
      return (
        <div
          onMouseDown={this.onMouseDown}
          className={this.props.classes.headlineButtonWrapper}
        >
          <button
            onClick={this.onClick}
            className={this.props.classes.headlineButton}
          >
            H
          </button>
        </div>
      );
    }
  },
);
const Editor: any = DraftEditor;
export const RichTextEditor = decorate<Props>(
  class RichTextEditorComponent extends React.Component<
    Props &
      WithStyles<
        | "editor"
        | "editorContainer"
        | "toolbar"
        | "headlineButton"
        | "headlineButtonWrapper"
      >,
    State
  > {
    state = {
      editorState:
        this.props.initialValue && this.props.initialValue !== ""
          ? EditorState.createWithContent(
              convertFromRaw(JSON.parse(this.props.initialValue)),
            )
          : EditorState.createEmpty(),
      focused: false,
    };

    editor: any = null;
    inlineToolbarPlugin = createInlineToolbarPlugin({
      structure: [
        BoldButton,
        ItalicButton,
        UnderlineButton,
        Separator,
        HeadlinesButton,
        UnorderedListButton,
        OrderedListButton,
        BlockquoteButton,
      ],
    });

    handleKeyCommand(command: any, editorState: any) {
      const newState = RichUtils.handleKeyCommand(editorState, command);
      if (newState) {
        this.onChange(newState);
        return "handled";
      }
      return "not-handled";
    }

    onBlur = () => this.setState({ focused: false });

    onChange = (editorState: any) => {
      const value = JSON.stringify(
        convertToRaw(editorState.getCurrentContent()),
      );
      const event = {
        target: { name: this.props.name, value },
      };
      this.props.onChange(event as any);
      this.setState({ editorState });
    };

    onFocus = () => this.setState({ focused: true });

    render() {
      const { classes, label, helperText } = this.props;
      const { InlineToolbar } = this.inlineToolbarPlugin;
      const plugins = [
        this.inlineToolbarPlugin,
        dragNDropFileUploadPlugin,
        blockDndPlugin,
        focusPlugin,
        alignmentPlugin,
        resizeablePlugin,
        imagePlugin,
      ];
      return (
        <div className={classes.editorContainer}>
          {label && <InputLabel label={label}>{null}</InputLabel>}
          <InputFocus focused={this.state.focused}>
            <div
              className={[
                classes.editor,
                this.state.focused ? "active" : undefined,
              ].join(" ")}
            >
              <Editor
                editorState={this.state.editorState}
                handleKeyCommand={this.handleKeyCommand}
                plugins={plugins}
                onBlur={this.onBlur}
                onChange={this.onChange}
                onFocus={this.onFocus}
              />
              <InlineToolbar className={{ toolbarStyles: classes.toolbar }} />
              {this.state.focused && <AlignmentTool />}
            </div>
          </InputFocus>
          <Typography variant="caption">
            {helperText ||
              i18n.t("Select text to enable text formatting tools")}
          </Typography>
        </div>
      );
    }
  },
);
export default RichTextEditor;
