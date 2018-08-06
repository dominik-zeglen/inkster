import * as React from "react";
import axios from "axios";

interface State {
  active: boolean;
  progress: number;
}
interface Props {
  children: ((props: State) => React.ReactElement<any>) | React.ReactNode;
}
interface FileUploadOptions {
  file: any;
  onError: () => void;
  onSuccess: (filename: string) => void;
}
interface Context {
  uploadFile: (opts: FileUploadOptions) => void;
}
export const {
  Provider: UploadContextProvider,
  Consumer: WithUpload
} = React.createContext<{
  uploadFile: (opts: FileUploadOptions) => void;
}>({
  uploadFile: () => {}
} as Context);

export class UploadProvider extends React.Component<Props, State> {
  state = {
    active: false,
    progress: 0
  };

  handleUpload = (opts: FileUploadOptions) => {
    this.setState({ active: true });
    const form: any = new FormData();
    form.append("file", opts.file);

    axios
      .post("/upload", form, {
        onUploadProgress: progressEvent => {
          const totalLength = progressEvent.lengthComputable
            ? progressEvent.total
            : progressEvent.target.getResponseHeader("content-length") ||
              progressEvent.target.getResponseHeader(
                "x-decompressed-content-length"
              );
          if (totalLength !== null) {
            this.setState({
              progress: Math.round((progressEvent.loaded * 100) / totalLength)
            });
          }
        }
      })
      .then(res => {
        this.setState({ active: false });
        opts.onSuccess(res.data.filename);
      })
      .catch(err => {
        this.setState({ active: false });
        opts.onError();
      });
  };

  render() {
    const { children } = this.props;
    return (
      <UploadContextProvider value={{ uploadFile: this.handleUpload }}>
        {typeof children === "function" ? children(this.state) : children}
      </UploadContextProvider>
    );
  }
}
export default UploadProvider;
