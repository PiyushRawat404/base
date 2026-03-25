
import styled from "styled-components";


const Text = styled.p`
  color: ${(props) => props.color};
  font-weight: ${(props) => (props.$bold ? "bold" : "normal")};
  font-style: ${(props) => (props.$italic ? "italic" : "normal")};
  text-decoration: ${(props) => (props.$underline ? "underline" : "none")};
  font-size: 20px;
`;

function App() {
  return (
    <div style={{ padding: "20px" }}>
      <Text color="red" $bold>Sample Text</Text>
      <Text color="blue" $italic>Sample Text</Text>
      <Text color="green" $underline>Sample TExt</Text>
    </div>
  );
}

export default App;