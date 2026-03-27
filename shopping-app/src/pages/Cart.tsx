import { useSelector } from "react-redux";
import type  { RootState } from "../store/store.ts";

const Cart = () => {
  const items = useSelector((state: RootState) => state.cart.items);

  return (
    <div>
      <h2>Cart</h2>

      {items.length === 0 ? (
        <p>No items in cart</p>
      ) : (
        items.map((item, index) => (
          <div key={index}>
            {item.name} - ₹{item.price}
          </div>
        ))
      )}
    </div>
  );
};

export default Cart;