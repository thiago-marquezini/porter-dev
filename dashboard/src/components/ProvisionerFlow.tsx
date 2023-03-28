import React, { useState, useContext, useMemo } from "react";
import styled from "styled-components";

import { integrationList } from "shared/common";
import { Context } from "shared/Context";

import ProvisionerForm from "components/ProvisionerForm";
import CredentialsForm from "components/CredentialsForm";
import Helper from "components/form-components/Helper";
import Modal from "./porter/Modal";
import Text from "./porter/Text";
import Spacer from "./porter/Spacer";
import Fieldset from "./porter/Fieldset";
import Checkbox from "./porter/Checkbox";
import Button from "./porter/Button";
import ExpandableSection from "./porter/ExpandableSection";
import Input from "./porter/Input";

const providers = ["aws", "gcp", "azure"];

type Props = {
};

const ProvisionerFlow: React.FC<Props> = ({
}) => {
  const { usage, hasBillingEnabled } = useContext(Context);
  const [currentStep, setCurrentStep] = useState("cloud");
  const [credentialId, setCredentialId] = useState("");
  const [showCostConfirmModal, setShowCostConfirmModal] = useState(false);
  const [confirmCost, setConfirmCost] = useState("");

  const isUsageExceeded = useMemo(() => {
    if (!hasBillingEnabled) {
      return false;
    }
    return usage?.current.clusters >= usage?.limit.clusters;
  }, [usage]);

  if (currentStep === "cloud") {
    return (
      <>
        <StyledProvisionerFlow>
          <Helper>
            Select your hosting backend:
          </Helper>
          <BlockList>
            {providers.map((provider: string, i: number) => {
              let providerInfo = integrationList[provider];
              return (
                <Block
                  key={i}
                  disabled={isUsageExceeded || provider === "gcp" || provider === "azure"}
                  onClick={() => {
                    if (!(isUsageExceeded || provider === "gcp" || provider === "azure")) {
                      setShowCostConfirmModal(true);
                    }
                  }}
                >
                  <Icon src={providerInfo.icon} />
                  <BlockTitle>{providerInfo.label}</BlockTitle>
                  <BlockDescription>{providerInfo.tagline || "Hosted in your own cloud"}</BlockDescription>
                </Block>
              );
            })}
          </BlockList>
        </StyledProvisionerFlow>
        {showCostConfirmModal && (
          <Modal closeModal={() => {
            setConfirmCost("");
            setShowCostConfirmModal(false);
          }}>
            <Text size={16} weight={500}>
              AWS base cost consent
            </Text>
            <Spacer height="15px" />
            <Text color="helper">
              Porter will create resources in your existing AWS account for hosting your applications. You will be separately charged by AWS and can use your cloud credits.
            </Text>
            <Spacer y={1} />
            <Text color="helper">
              AWS base cost before cloud credits:
            </Text>
            <Spacer y={1} />
            <ExpandableSection
              background="#ffffff11"
              Header={
                <Cost>$315.94 / mo</Cost>
              }
              ExpandedSection={
                <Dark>
                  <Text>
                    • Amazon Elastic Kubernetes Service (EKS) = $73/mo
                    <Spacer height="15px" />
                    • Amazon EC2:
                    <Spacer height="15px" />
                    <Tab />+ System workloads: t3.medium instance (2) = $60.74/mo
                    <Spacer height="15px" />
                    <Tab />+ Monitoring workloads: t3.large instance (1) = $60.74/mo
                    <Spacer height="15px" />
                    <Tab />+ Application workloads: t3.xlarge instance (1) = $121.47/mo
                  </Text>
                </Dark>
              }
            />
            <Spacer y={1} />
            <Text color="helper">
              Porter metered cost: $0.019/hr/vCPU + $0.009/hr/GB RAM.
            </Text>
            <Spacer y={1} />
            <Text color="helper">
              All AWS resources will be automatically deleted when you delete your Porter project. Please enter the base cost ("315.94") below to proceed:
            </Text>
            <Spacer y={1} />
            <Input placeholder="315.94" value={confirmCost} setValue={setConfirmCost} width="100%" height="40px" />
            <Spacer y={1} />
            <Button
              disabled={confirmCost !== "315.94"}
              onClick={() => {
                setShowCostConfirmModal(false);
                setConfirmCost("");
                setCurrentStep("credentials");
              }}
            >
              Continue
            </Button>
          </Modal>
        )}
      </>
    );
  } else if (currentStep === "credentials") {
    return (
      <CredentialsForm 
        goBack={() => setCurrentStep("cloud")}
        proceed={(id) => {
          setCredentialId(id);
          setCurrentStep("cluster");
        }}
      />
    );
  } else if (currentStep === "cluster") {
    return (
      <ProvisionerForm
        goBack={() => setCurrentStep("credentials")}
        credentialId={credentialId}
      />
    );
  }
};

export default ProvisionerFlow;

const Dark = styled.div`
  position: relative;
  padding: 25px;
  background: #1b1d2688;
  font-size: 13px;
`;

const Cost = styled.div`
  font-weight: 600;
  font-size: 20px;
`;

const Tab = styled.span`
  margin-left: 20px;
  height: 1px;
`;

const BlockList = styled.div`
  overflow: visible;
  margin-top: 25px;
  margin-bottom: 27px;
  display: grid;
  grid-column-gap: 25px;
  grid-row-gap: 25px;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
`;

const Icon = styled.img<{ bw?: boolean }>`
  height: 42px;
  margin-top: 30px;
  margin-bottom: 15px;
  filter: ${(props) => (props.bw ? "grayscale(1)" : "")};
`;

const BlockDescription = styled.div`
  margin-bottom: 12px;
  color: #ffffff66;
  text-align: center;
  font-weight: 400;
  font-size: 13px;
  padding: 0px 25px;
  height: 2.4em;
  font-size: 12px;
  display: -webkit-box;
  overflow: hidden;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
`;

const BlockTitle = styled.div`
  margin-bottom: 12px;
  width: 80%;
  text-align: center;
  font-size: 14px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
`;

const Block = styled.div<{ disabled?: boolean }>`
  align-items: center;
  user-select: none;
  display: flex;
  font-size: 13px;
  overflow: hidden;
  font-weight: 500;
  padding: 3px 0px 5px;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  height: 170px;
  filter: ${({ disabled }) => (disabled ? "brightness(0.8) grayscale(1)" : "")};
  cursor: ${(props) => (props.disabled ? "not-allowed" : "pointer")};
  color: #ffffff;
  position: relative;
  border-radius: 5px;
  background: #26292e;
  border: 1px solid #494b4f;
  :hover {
    border: ${(props) => (props.disabled ? "" : "1px solid #7a7b80")};
  }

  animation: fadeIn 0.3s 0s;
  @keyframes fadeIn {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }
`;

const StyledProvisionerFlow = styled.div`
  margin-top: -24px;
`;