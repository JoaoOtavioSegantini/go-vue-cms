import { mount } from "@vue/test-utils";
import EditPage from "@/components/Users/Edit.vue";
import ClassicEditor from "@ckeditor/ckeditor5-build-classic";
import { createStore } from "vuex";

describe("EditUser.vue", () => {
  afterEach(() => {
    jest.restoreAllMocks();
  });

  window.axios = {
    get: jest
      .fn(() => {
        return {
          then: jest.fn(() => "your faked response"),
        };
      })
      .mockName("axiosget"),
    put: jest
      .fn(() => {
        return {
          then: jest.fn(() => "your faked response"),
        };
      })
      .mockName("axiosPut"),
  };

  const mockRoute = {
    params: {
      id: 1,
    },
  };
  const mockRouter = {
    push: jest.fn(),
  };
  const createVuexStore = () => {
    return createStore({
      state() {
        return {
          Users: {
            oneUser: {
              ID: "1",
              name: "Post 1",
              email: "fake@gmail.com",
              username: "John Doe",
              password: "123",
            },
          },
        };
      },
      mutations: {
        updateUser(state) {
          state.oneUser = {
            ID: "1",
            slug: "John Doe",
            title: "Post 1",
            body: "A simple post text body",
          };
        },
      },
      actions: {
        getUser(ctx, res) {
          ctx.commit("updateUser", res);
        },
        updateUser(ctx, data) {},
      },
    });
  };

  function factory() {
    const store = createVuexStore();

    return mount(EditPage, {
      mocks: {
        axios: window.axios,
      },
      global: {
        plugins: [store],
        mocks: {
          $router: mockRouter,
          $route: mockRoute,
        },
      },
      data() {
        return {
          editor: ClassicEditor,
        };
      },
    });
  }

  it("renders props.msg when passed", async () => {
    const msg =
      "AÇÕES acbdfg@gmail.com Edição de usuário do blogNomeUsernameEmailSenhaSalvar";
    const wrapper = factory();

    wrapper.findAll("input")[0].setValue("acbdfg@gmail.com");
    wrapper.findAll("input")[1].setValue("987654");

    await wrapper.find("form").trigger("submit");
    await wrapper.find("button").trigger("click");
    await wrapper.vm.save();
    expect(mockRouter.push).toBeCalledWith({ path: "/users/1" });
    expect(wrapper.text()).toMatch(msg);
    expect(wrapper.html()).toMatchSnapshot();
  });
});
