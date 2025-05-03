export interface EmployeeType {
  employee_id: number;
  employee_username: string;
  employee_uuid: string;
  employee_position: number;
  employee_full_name: number;
  employee_phone_number: string;
  employee_email_address: string;
  created_at: string;
  created_by_employee_id: number;
  employee_role: string[];
  is_banned: boolean;
  is_suspended: boolean;
  suspension_duration: any;
}
