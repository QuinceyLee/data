package net.sppan.base.dao;

import org.hibernate.query.NativeQuery;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import net.sppan.base.dao.support.IBaseDao;
import net.sppan.base.entity.Role;

@Repository
public interface IRoleDao extends IBaseDao<Role, Integer> {

	Page<Role> findAllByNameContainingOrDescriptionContaining(String searchText1,String searchText2, Pageable pageable);

	@Modifying
	@Query(nativeQuery = true,value="delete from tb_user_role where role_id = :roleId")
	void deleteRoleUserRelation(@Param("roleId")Integer roleId);

	@Modifying
	@Query(nativeQuery = true,value="delete from tb_role_resource where role_id = :roleId")
	void deleteRoleResourceRelation(@Param("roleId")Integer roleId);
}
