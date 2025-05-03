export interface PorductType {
  product_id: number;
  product_title: string;
  product_description: string;
  product_price: number;
  created_by_employee_id: number;
  product_cover: string;
  created_at: string;
}
export interface CreateProductType {
  product_title: string;
  product_price: any;
  product_description: string;
  product_cover: any;
}
export type CreateProductResponseType = {
    status: number,
    message: string
}
export type SelectedProduct = {
  product_id: number;
  product_title: string;
  product_description: string;
  product_price: number;
  created_by_employee_id: number;
  product_cover: string;
  created_at: string;
}