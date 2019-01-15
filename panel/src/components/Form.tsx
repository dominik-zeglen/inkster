import * as React from "react";

export type FormChildren<T extends {}> = ((
  props: {
    data: T;
    hasChanged: boolean;
    change: (event: React.ChangeEvent<any>) => void;
    submit: (event: React.FormEvent<any>) => void;
  },
) => React.ReactElement<any>);

export interface FormProps<T extends {}> {
  children: FormChildren<T>;
  initial: T;
  onSubmit: (data?: T) => void;
}

const shallowCompare = (a: any, b: any) => {
  let ret = true;
  Object.keys(a).forEach(k => {
    if (typeof a[k] === "object") {
      if (JSON.stringify(a[k]) !== JSON.stringify(b[k])) {
        ret = false;
      }
    } else {
      if (a[k] !== b[k]) {
        ret = false;
      }
    }
  });
  return ret;
};

class Form<T extends {} = {}> extends React.Component<FormProps<T>, T> {
  public state: T = this.props.initial;
  private form = React.createRef<HTMLFormElement>();

  private handleChange = (event: React.ChangeEvent<any>) => {
    const { target } = event;
    if (!(target.name in this.state)) {
      console.error(`Unknown form field: ${target.name}`);
      return;
    }
    this.setState(({ [target.name]: target.value } as any) as Pick<T, keyof T>);
  };

  private handleSubmit = (event: React.FormEvent<any>) => {
    const { onSubmit } = this.props;
    event.preventDefault();
    if (!!this.form.current) {
      if (this.form.current.checkValidity()) {
        onSubmit(this.state);
      } else {
        (this.form.current.querySelectorAll("input") as any).forEach(input =>
          input.reportValidity(),
        );
      }
    }
  };

  public render() {
    const { children } = this.props;
    return (
      <form ref={this.form} onSubmit={this.handleSubmit}>
        {children({
          change: this.handleChange,
          data: this.state,
          hasChanged: !shallowCompare(this.props.initial, this.state),
          submit: this.handleSubmit,
        })}
      </form>
    );
  }
}

export default Form;
