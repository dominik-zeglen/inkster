import * as React from "react";
import { Image as ImageIcon, File as FileIcon, X } from "react-feather";
import Button from "aurora-ui-kit/dist/components/Button";
import Card from "aurora-ui-kit/dist/components/Card";
import CardContent from "aurora-ui-kit/dist/components/CardContent";
import CardHeader from "aurora-ui-kit/dist/components/CardHeader";
import CardTitle from "aurora-ui-kit/dist/components/CardTitle";
import IconButton from "aurora-ui-kit/dist/components/IconButton";
import Input from "aurora-ui-kit/dist/components/TextInput";
import InputLabel from "aurora-ui-kit/dist/components/InputLabel";

import RichTextEditor from "../../components/RichTextEditor";
import i18n from "../../i18n";
import Spacer from "../../components/Spacer";

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
    cb: (event: React.ChangeEvent<any>) => void,
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
    <Card {...props}>
      <CardHeader>
        <CardTitle>{i18n.t("Field properties")}</CardTitle>
        <IconButton onClick={onDelete}>
          <X />
        </IconButton>
      </CardHeader>
      <CardContent>
        <Input
          label={i18n.t("Name")}
          value={data.name}
          onChange={value =>
            onChange({
              target: {
                name: "name",
                value,
              },
            } as any)
          }
        />
        <Spacer />
        {data.type === "longText" ? (
          <RichTextEditor
            label={i18n.t("Field value")}
            initialValue={data.value}
            name="value"
            onChange={onChange}
          />
        ) : data.type === "image" ? (
          <div
            style={{
              paddingTop: 16,
            }}
          >
            <InputLabel label={i18n.t("Image")}>
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
            </InputLabel>
            <div>
              {data.value ? (
                <img
                  src={"/static/" + data.value}
                  style={{
                    borderRadius: 6,
                    height: "auto",
                    display: "block",
                    maxWidth: "100%",
                    margin: "10px 0",
                  }}
                />
              ) : (
                <span
                  style={{
                    display: "block" as "block",
                    margin: "20px auto",
                    width: 64,
                  }}
                >
                  <ImageIcon size={64} />
                </span>
              )}
              <Button
                onClick={
                  this.refs ? () => this.refs[data.name].click() : undefined
                }
                componentProps={{
                  style: { width: "100%" },
                }}
                variant="outlined"
                color="secondary"
              >
                {i18n.t("Upload image")}
              </Button>
            </div>
          </div>
        ) : data.type === "file" ? (
          <>
            <InputLabel label={i18n.t("File")}>
              <input
                name="value"
                type="file"
                ref={ref => {
                  this.refs[data.name] = ref;
                }}
                style={{ display: "none" as "none" }}
                onChange={onUpload(onChange)}
              />
            </InputLabel>
            <div>
              {data.value ? (
                <a href={"/static/" + data.value}>
                  {i18n.t("Download {{ filename }}", {
                    filename: data.value,
                  })}
                </a>
              ) : (
                <span
                  style={{
                    display: "block" as "block",
                    margin: "20px auto",
                    width: 64,
                  }}
                >
                  <FileIcon size={64} />
                </span>
              )}
              <Button
                onClick={
                  this.refs ? () => this.refs[data.name].click() : undefined
                }
                componentProps={{
                  style: { width: "100%" },
                }}
                variant="outlined"
                color="secondary"
              >
                {i18n.t("Change file")}
              </Button>
            </div>
          </>
        ) : (
          <Input
            label={i18n.t("Value")}
            value={data.value}
            onChange={value =>
              onChange({
                target: {
                  name: "value",
                  value,
                },
              } as any)
            }
          />
        )}
      </CardContent>
    </Card>
  );
};

export default PageFieldProperties;
