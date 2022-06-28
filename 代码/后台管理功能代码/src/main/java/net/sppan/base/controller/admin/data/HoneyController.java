package net.sppan.base.controller.admin.data;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

import net.sppan.base.controller.BaseController;
import net.sppan.base.service.IRoleService;
import net.sppan.base.service.IUserService;
import org.springframework.web.bind.annotation.ResponseBody;

@Controller
@RequestMapping("/admin/honeyed")
public class HoneyController extends BaseController {

    @Autowired
    private IUserService userService;
    @Autowired
    private IRoleService roleService;

    @RequestMapping(value = { "/", "/index" })
    public String index() {
        return "admin/honeyed/index";
    }
    @ResponseBody
    @RequestMapping(value = { "/test" })
    public hhh index1() {
        hhh h = new hhh();
        h.setCode("403");
        h.setMessage("权限不足");
        return h;
    }





}
