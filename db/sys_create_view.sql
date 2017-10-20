create or replace VIEW
V_Sys_FamilyUserMembers
AS
select 
	t3.ID,
	t3.UserID,
	t3.FamilyID,
	t3.UserStatus,
	t1.OpenID,
	t1.`NAME` UserName,
	t1.Photo UserPhoto,
	t1.Description UserDescription,
	t2.`NAME` FamilyName,
	t2.Photo FamilyPhoto,
	t2.Description FamilyDescription,
	t2.Label FamilyLabel,
	t2.CreateBy FamilyCreateBy,
	t2.CreateTime FamilyCreateTime
from sys_user t1,sys_family t2,sys_familymember t3
	where t1.ID=t3.UserID and t2.ID=t3.FamilyID