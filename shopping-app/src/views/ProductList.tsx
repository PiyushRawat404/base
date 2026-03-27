import { products } from "../lib/products.ts";
import ProductCard from "../components/ProductCard";

const ProductList = () => {
  return (
    <div>
      <h2>Products</h2>
      {products.map((p) => (
        <ProductCard key={p.id} product={p} />
      ))}
    </div>
  );
};

export default ProductList;