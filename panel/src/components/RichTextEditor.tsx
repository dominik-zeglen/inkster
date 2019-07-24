import classNames from "classnames";
import * as React from "react";
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
  AlignBlockDefaultButton,
  AlignBlockLeftButton,
  AlignBlockCenterButton,
  AlignBlockRightButton,
} from "draft-js-buttons";
import createImagePlugin from "draft-js-image-plugin";
import createFocusPlugin from "draft-js-focus-plugin";
import createResizeablePlugin from "draft-js-resizeable-plugin";
import createBlockDndPlugin from "draft-js-drag-n-drop-plugin";
import createDragNDropUploadPlugin from "@mikeljames/draft-js-drag-n-drop-upload-plugin";
import InputFocus from "aurora-ui-kit/dist/components/InputFocus";
import InputLabel from "aurora-ui-kit/dist/components/InputLabel";
import Typography from "aurora-ui-kit/dist/components/Typography";
import createUseStyles, { css } from "aurora-ui-kit/dist/utils/jss";
import { ITheme } from "aurora-ui-kit/dist/theme";

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

const useStyles = createUseStyles((theme: ITheme) => ({
  button: {
    "&": css`
      -webkit-appearance: none;
      align-items: center;
      background: ${theme.mixins.fade(
        theme.colors.primary.main,
        theme.alpha.light,
      )};
      border: none;
      color: ${theme.colors.gray.dark};
      cursor: pointer;
      display: flex;
      height: 32px;
      justify-content: center;
      padding: 0;
      width: 32px;
    `,
    ":focus, &:hover": css`
      background: ${theme.mixins.fade(
        theme.colors.primary.main,
        theme.alpha.default,
      )};
    `,
  },
  buttonActive: {
    "&$button": css`
      background: ${theme.mixins.fade(
        theme.colors.primary.main,
        theme.alpha.default,
      )};
    `,
  },
  editor: {
    background: theme.mixins.fade(theme.colors.primary.main, theme.alpha.light),
    border: `1px solid ${theme.colors.primary.lightest}`,
    marginBottom: theme.spacing,
    marginTop: 0,
    padding: theme.spacing,
    transition: theme.transition.default,
  },
  editorContainer: {
    paddingTop: theme.spacing * 3,
  },
  headlineButton: css`
    border: 0;
    color: ${theme.colors.common.black};
    font-size: 21px;
    position: relative;
    vertical-align: bottom;
  `,
  headlineButtonWrapper: {
    display: "inline-block" as "inline-block",
  },
  toolbar: css`
    background: ${theme.colors.background.main};
    border: 1px solid
      ${theme.mixins.fade(theme.colors.gray.main, theme.alpha.default)};
    border-radius: 2px;
    display: flex;
    padding: ${theme.spacing}px;
    position: absolute;
  `,
}));

const AlignmentPicker: React.FC<any> = ({ onOverrideContent, ...props }) => {
  const classes = useStyles();

  const onWindowClick = () => onOverrideContent(undefined);

  React.useEffect(() => {
    window.addEventListener("click", onWindowClick);

    return () => window.removeEventListener("click", onWindowClick);
  });

  const buttons = [
    AlignBlockDefaultButton,
    AlignBlockLeftButton,
    AlignBlockCenterButton,
    AlignBlockRightButton,
  ];

  return (
    <div className={classes.toolbar}>
      {buttons.map((Button, i) => (
        <Button key={i} {...props} />
      ))}
    </div>
  );
};
const AlignmentButton: React.FC<any> = ({ onOverrideContent }) => {
  const classes = useStyles();

  return (
    <div
      onMouseDown={event => event.preventDefault()}
      className={classes.headlineButtonWrapper}
    >
      <button
        onClick={() => onOverrideContent(AlignmentPicker)}
        className={classNames(classes.button, classes.headlineButton)}
      >
        A
      </button>
    </div>
  );
};

const HeadlinesPicker: React.FC<any> = ({ onOverrideContent, ...props }) => {
  const classes = useStyles();

  const onWindowClick = () => onOverrideContent(undefined);

  React.useEffect(() => {
    window.addEventListener("click", onWindowClick);

    return () => window.removeEventListener("click", onWindowClick);
  });

  const buttons = [HeadlineOneButton, HeadlineTwoButton, HeadlineThreeButton];

  return (
    <div className={classes.toolbar}>
      {buttons.map((Button, i) => (
        <Button key={i} {...props} />
      ))}
    </div>
  );
};
const HeadlinesButton: React.FC<any> = ({ onOverrideContent }) => {
  const classes = useStyles();

  return (
    <div
      onMouseDown={event => event.preventDefault()}
      className={classes.headlineButtonWrapper}
    >
      <button
        onClick={() => onOverrideContent(HeadlinesPicker)}
        className={classNames(classes.button, classes.headlineButton)}
      >
        H
      </button>
    </div>
  );
};
const Editor: any = DraftEditor;

class RichTextEditorComponent extends React.Component<
  Props & {
    classes: Record<
      | "button"
      | "buttonActive"
      | "editor"
      | "editorContainer"
      | "toolbar"
      | "headlineButton"
      | "headlineButtonWrapper",
      string
    >;
  },
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
      AlignmentButton,
    ],
    theme: {
      buttonStyles: {
        active: this.props.classes.buttonActive,
        button: this.props.classes.button,
      },
      toolbarStyles: {
        toolbar: this.props.classes.toolbar,
      },
    },
  });

  decorator = composeDecorators(
    resizeablePlugin.decorator,
    focusPlugin.decorator,
    blockDndPlugin.decorator,
  );

  imagePlugin = createImagePlugin({
    decorator: this.decorator,
    theme: {
      buttonStyles: {
        active: this.props.classes.buttonActive,
        button: this.props.classes.button,
      },
      toolbarStyles: {
        toolbar: this.props.classes.toolbar,
      },
    },
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
    const value = JSON.stringify(convertToRaw(editorState.getCurrentContent()));
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
    const dragNDropFileUploadPlugin = createDragNDropUploadPlugin({
      handleUpload: mockUpload,
      addImage: this.imagePlugin.addImage,
    });
    const plugins = [
      this.inlineToolbarPlugin,
      dragNDropFileUploadPlugin,
      blockDndPlugin,
      focusPlugin,
      resizeablePlugin,
      this.imagePlugin,
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
            <InlineToolbar />
          </div>
        </InputFocus>
        <Typography variant="caption">
          {helperText || i18n.t("Select text to enable text formatting tools")}
        </Typography>
      </div>
    );
  }
}

export const RichTextEditor: React.FC<Props> = props => {
  const classes = useStyles();
  return <RichTextEditorComponent {...props} classes={classes} />;
};
export default RichTextEditor;
