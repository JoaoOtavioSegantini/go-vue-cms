import { mount } from "@vue/test-utils";
import AddPage from "@/components/Pages/Add.vue";
import ClassicEditor from "@ckeditor/ckeditor5-build-classic";
import { createStore } from "vuex";

 window.axios = {
    post: jest
      .fn(()  => {
        return {
          then: jest.fn(() => "your faked response"),
        };
      })
      .mockName("axiosPost"),
  };

describe("AddPage.vue", () => {
  afterEach(() => {
    jest.restoreAllMocks();
  });

  const mockRouter = {
    push: jest.fn(),
  };
  const createVuexStore = () => {
    return createStore({
      state() {
        return {
          all: [{ ID: "1", slug: "John Doe", title: "johndoe@gmail.com" }],
        };
      },
      mutations: {
        addToPagesList(state, res) {
          state.all.push(res);
        },
      },
      actions: {
        createPage(ctx, res) {
          ctx.commit("addToPagesList", res);
        },
      },
    });
  };

  function factory() {
    const store = createVuexStore();

    return mount(AddPage, {
        stubs: {
            ckeditor: true
          },
      
      mocks: {
        axios: window.axios,
      },
      global: {
        plugins: [store],
        mocks: {
          $router: mockRouter,
        //   $route: mockRoute,
        },
      },
      data() {
        return {
          editor: ClassicEditor,
          page: { ID: "1", title: "page 1" },
          editorData: "<p>Content of the editor.</p>",
        };
      },
    });
  }

  it("renders props.msg when passed", async () => {
    const msg =
      "AÇÕES Nova página Inclusão de nova página no siteTítuloURLConteúdoSalvar";
    const wrapper = factory();

    wrapper.findAll("input")[0].setValue("acbdfg@gmail.com");
    wrapper.findAll("input")[1].setValue("987654");

    await wrapper.find("form").trigger("submit");
    await wrapper.find("button").trigger("click");

    await wrapper.vm.save()
    console.log(wrapper.emitted())

    //expect(window.axios.post).toBeCalledWith({});
    expect(mockRouter.push).toBeCalledWith({ "path": "/pages"});
    expect(wrapper.text()).toMatch(msg);
    expect(wrapper.html()).toMatchSnapshot();
  });
});
