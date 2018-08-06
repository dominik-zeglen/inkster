import * as React from "react";
import { Button, Image, ControlLabel, Panel } from "react-bootstrap";
import { Image as ImageIcon, File as FileIcon, X } from "react-feather";

import IconButton from "../../components/IconButton";
import Input from "../../components/Input";
import RichTextEditor from "../../components/RichTextEditor";
import i18n from "../../i18n";

interface Props {
  data: {
    id: string;
    name: string;
    type: string;
    value: string;
  };
  name: string;
  onChange: (event: React.ChangeEvent<any>) => void;
  onDelete: () => void;
  onUpload: (
    cb: (event: React.ChangeEvent<any>) => void
  ) => (event: React.ChangeEvent<any>) => void;
}

export const PageFieldProperties: React.StatelessComponent<Props> = ({
  data,
  name,
  onChange,
  onDelete,
  onUpload,
  ...props
}) => {
  this.refs = {};
  return (
    <Panel {...props}>
      <Panel.Heading>
        <Panel.Title>{i18n.t("Field properties")}</Panel.Title>
        <IconButton icon={X} onClick={onDelete} />
      </Panel.Heading>
      <Panel.Body>
        <Input
          label={i18n.t("Name")}
          name="name"
          value={data.name}
          onChange={onChange}
        />
        {data.type === "longText" ? (
          <RichTextEditor
            label={i18n.t("Field value")}
            initialValue={data.value}
            name="value"
            onChange={onChange}
          />
        ) : data.type === "image" ? (
          <>
            <ControlLabel>{i18n.t("Image")}</ControlLabel>
            <input
              name="value"
              type="file"
              ref={ref => {
                this.refs[data.name] = ref;
              }}
              accept="image/*"
              style={{ display: "none" as "none" }}
              onChange={onUpload(onChange)}
            />
            <div>
              {data.value ? (
                <>
                  <Image
                    src={"/static/" + data.value}
                    rounded={true}
                    responsive={true}
                    style={{ margin: "10px 0" }}
                  />
                  <Button
                    onClick={
                      this.refs ? () => this.refs[data.name].click() : undefined
                    }
                    style={{ width: "100%" }}
                  >
                    {i18n.t("Change image")}
                  </Button>
                </>
              ) : (
                <>
                  <span
                    style={{
                      display: "block" as "block",
                      margin: "20px auto",
                      width: 64
                    }}
                  >
                    <ImageIcon size={64} />
                  </span>
                  <Button
                    onClick={
                      this.refs ? () => this.refs[data.name].click() : undefined
                    }
                    style={{ width: "100%" }}
                  >
                    {i18n.t("Upload image")}
                  </Button>
                </>
              )}
            </div>
          </>
        ) : data.type === "file" ? (
          <>
            <ControlLabel>{i18n.t("File")}</ControlLabel>
            <input
              name="value"
              type="file"
              ref={ref => {
                this.refs[data.name] = ref;
              }}
              style={{ display: "none" as "none" }}
              onChange={onUpload(onChange)}
            />
            <div>
              {data.value ? (
                <>
                  <a href={"/static/" + data.value}>
                    {i18n.t("Download {{ filename }}", {
                      filename: data.value
                    })}
                  </a>
                  <Button
                    onClick={
                      this.refs ? () => this.refs[data.name].click() : undefined
                    }
                    style={{ width: "100%" }}
                  >
                    {i18n.t("Change file")}
                  </Button>
                </>
              ) : (
                <>
                  <span
                    style={{
                      display: "block" as "block",
                      margin: "20px auto",
                      width: 64
                    }}
                  >
                    <FileIcon size={64} />
                  </span>
                  <Button
                    onClick={
                      this.refs ? () => this.refs[data.name].click() : undefined
                    }
                    style={{ width: "100%" }}
                  >
                    {i18n.t("Upload file")}
                  </Button>
                </>
              )}
            </div>
          </>
        ) : (
          <Input
            label={i18n.t("Value")}
            name="value"
            value={data.value}
            onChange={onChange}
          />
        )}
      </Panel.Body>
    </Panel>
  );
};

export default PageFieldProperties;
