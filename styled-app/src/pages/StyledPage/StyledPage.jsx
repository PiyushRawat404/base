
import styled from "styled-components";
import "../../styles/StyledPage.css";

const Text = styled.p`
  color: ${(props) => props.color};
  font-weight: ${(props) => (props.$bold ? "bold" : "normal")};
  font-style: ${(props) => (props.$italic ? "italic" : "normal")};
  text-decoration: ${(props) => (props.$underline ? "underline" : "none")};
  font-size: 20px;
`;

const StyledPage = () => {
  return (
    <div className="styledcontainer" style={{ padding: "20px" }}>
      <h2>Styled Components Page</h2>

      <Text color="red" $bold>
        Sample Text
      </Text>

      <Text color="blue" $italic>
        Sample Text
      </Text>

      <Text color="green" $underline>
        Sample Text
      </Text>
    </div>
  );
};

export default StyledPage;