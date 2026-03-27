import { useDispatch } from "react-redux";
import { addToCart } from "../features/cartSlice";

type Props = {
  product: {
    id: number;
    name: string;
    price: number;
  };
};

const ProductCard = ({ product }: Props) => {
  const dispatch = useDispatch();

  return (
    <div style={{ border: "1px solid black", padding: "10px", margin: "10px" }}>
      <h3>{product.name}</h3>
      <p>₹{product.price}</p>

      <button onClick={() => dispatch(addToCart(product))}>
        Add to Cart
      </button>
    </div>
  );
};

export default ProductCard;