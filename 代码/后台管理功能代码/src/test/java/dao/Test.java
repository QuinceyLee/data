package dao;

import net.sppan.base.Application;
import net.sppan.base.dao.IAttackMsgDao;
import net.sppan.base.dao.IRoleDao;
import net.sppan.base.entity.AttackedMsg;
import net.sppan.base.entity.Role;
import net.sppan.base.service.IAttackedMsgService;
import net.sppan.base.service.IResourceService;
import net.sppan.base.service.IRoleService;
import net.sppan.base.service.IUserService;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;

import javax.annotation.Resource;
import java.util.List;

@SpringBootTest(classes = Application.class)
@RunWith(SpringJUnit4ClassRunner.class)
public class Test
{

    @Resource
    IRoleService iRoleService;

    @Resource
    IUserService iUserService;

    @Resource
    IResourceService iResourceService;

    @Resource
    IAttackedMsgService iAttackedMsgService;

    @org.junit.Test
    public void test(){
        Page<AttackedMsg> all = iAttackedMsgService.findAll(new PageRequest(0, 2));
        System.out.println(all.getContent().get(0));
    }
}
