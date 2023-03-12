import { mount } from "@vue/test-utils";
import EditPage from "@/components/Posts/Edit.vue";
import ClassicEditor from "@ckeditor/ckeditor5-build-classic";
import { createStore } from "vuex";

describe("EditPost.vue", () => {
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
          Posts: {
            onePost: {
              ID: "1",
              title: "Post 1",
              body: "A simple post text body",
              slug: "simple-slug-text",
            },
          },
        };
      },
      mutations: {
        updatePost(state) {
          state.onePost = {
            ID: "1",
            slug: "John Doe",
            title: "Post 1",
            body: "A simple post text body",
          };
        },
      },
      actions: {
        getPost(ctx, res) {
          ctx.commit("updatePost", res);
        },
        updatePost(ctx, data) {},
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
    const msg = "AÇÕES Post 1 Edição de artigo do blogTítuloURLConteúdoSalvar";
    const wrapper = factory();

    wrapper.findAll("input")[0].setValue("acbdfg@gmail.com");
    wrapper.findAll("input")[1].setValue("987654");

    await wrapper.find("form").trigger("submit");
    await wrapper.find("button").trigger("click");
    await wrapper.vm.save();
    expect(mockRouter.push).toBeCalledWith({ path: "/posts/1" });
    expect(wrapper.text()).toMatch(msg);
    expect(wrapper.html()).toMatchSnapshot();
  });
});
